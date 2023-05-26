package loader

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	runtimeDebug "runtime/debug"
	"time"

	"github.com/mowshon/moviego"
)

var (
	// лог информации для отладки. По-умолчанию весь вывод этого лога идёт в io.Discard, что аналогично направлению в /dev/null
	debl                                               = log.New(io.Discard, "[DEBUG]\t", log.Ldate|log.Ltime|log.Lshortfile)
	errl                                               = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	debug, help, notDeleteTempFile, overwrite, verMode bool
	tmp, baseURL, output, useragent                    string
	timeout                                            int
)

const defaultUserAgent = `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:98.0) Gecko/20100101 Firefox/98.0`

func init() {
	flag.BoolVar(&help, "help", false, "shows help; equals `help` command; ignores other flags")
	flag.StringVar(&output, "output", "output.gif", "output file; output.gif")
	flag.StringVar(&output, "o", "..//output/output.gif", "alias for --output")
	flag.StringVar(&tmp, "tmp", "..//temp", "temp directory")
	flag.BoolVar(&notDeleteTempFile, "not-del-temp", true, "doesn't delete temp file if put")
	flag.BoolVar(&overwrite, "overwrite", true, "overwrites output file if it exists")
	flag.StringVar(&baseURL, "url", "https://randomcatgifs.com/", "url of site (idk why)")
	flag.StringVar(&useragent, "useragent", defaultUserAgent, "User-Agent header content")
	flag.IntVar(&timeout, "timeout", 10, "count of seconds to get gifs")
	flag.IntVar(&timeout, "t", 10, "alias for --timeout")
	flag.BoolVar(&debug, "debug", false, "turns on debug log")
	flag.BoolVar(&verMode, "version", false, "prints version end exits")

	flag.Parse()

	if help || (len(flag.Args()) > 0 && flag.Args()[0] == "help") {
		fmt.Printf("Syntax: %s [flags]\n", os.Args[0])
		flag.Usage()
		os.Exit(0)
	}

	if verMode { // получаем версию модуля, в которой был сбилдена программа
		var version string
		if bInfo, ok := runtimeDebug.ReadBuildInfo(); ok && bInfo.Main.Version != "(devel)" {
			version = bInfo.Main.Version
		} else {
			version = "unknown/not-versioned build"
		}
		fmt.Println(version)
		os.Exit(0)

	}

	if len(flag.Args()) > 0 {
		errl.Println("have too much args.")
		fmt.Printf("Syntax: %s [flags]\n", os.Args[0])
		os.Exit(1)
	}

	if timeout <= 0 {
		errl.Println("timeout must be greater than zero")
		os.Exit(1)
	}

	if !debug {
		/*
		   по-умолчанию библиотека для работы с ffmpeg выводит свою
		   итоговую команду с помощью log (похоже, забыли удалить/закомментировать это)
		   поэтому мы убираем вывод log'а, если мы не хотим видеть отладочную информацию
		*/
		log.SetOutput(io.Discard)
	} else {
		debl.SetOutput(os.Stderr)
	}
}

func GetGif() {

	var client = NewClient(
		BaseURL(baseURL),
		TempDir(tmp),
		UserAgent(useragent),
	)
	client.Debug = debug
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	video, err := client.GetVideo(context)
	if err != nil {
		errl.Printf("%v\n", err)
		os.Exit(1)
	}
	vidpath, err := client.SaveVideoToTemp(video)
	if err != nil {
		errl.Printf("%v\n", err)
		os.Exit(1)
	}

	// Edit size video
	editvideo, err := moviego.Load(vidpath)
	if err != nil {
		log.Println("video don't load")
	}
	newVidPath := "..//temp/newvideo.webm"
	err = editvideo.Resize(250, 250).Output(newVidPath).Run()
	if err != nil {
		log.Println("video don't save")
	}

	err = client.Convert(newVidPath, output, overwrite)
	if err != nil {
		var addition string
		// такая ошибка часто появляется поскольку не установлена консольная команда ffmpeg (консольная утилита)
		if err.Error() == "exit status 1" {
			addition = ". (This error often happens when file already exists)"
		}
		errl.Printf("%v%s\n", err, addition)
		os.Exit(1)
	}
	if !notDeleteTempFile {
		os.Remove(newVidPath)
		err := os.Remove(vidpath)
		if err != nil {
			errl.Printf("%v\n", err)
			os.Exit(1)
		}
	}

	// выводим имя файла, в который сохраняем.
	// не то, чтобы это было сильно полезно,
	// но это будет приятным (или нет) бонусом, если
	// понадобится что-то делать с получившимся
	// файлом после сохранения

	fmt.Printf("%s\n", output)
}

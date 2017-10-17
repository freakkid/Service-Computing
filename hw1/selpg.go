package main

import (
	// "log"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ERROR_TYPE int

const (
	EMPTY_STRING = ""
)

// 错误类型的枚举
const (
	INVALID_S_NUM        ERROR_TYPE = 1 + iota // 1
	INVALID_E_NUM                              // 2
	INVALID_PAGE_LEN                           // 3
	NOT_EXIST_INPUT_FILE                       // 4
	NOREADALE_INPUT_FILE                       // 5
	NOT_OPEN_INPUT_FILE                        // 6
	NOT_OPEN_PIPE                              // 7
	NOT_START_PRINTER						   // 8
	S_PAGE_GREATER_TOTAL_PAGES 				   // 9
	E_PAGE_GREATER_TOTAL_PAGES 				   // 10
	COMPLETE_PRINT_ERROR					   // 11
	OTHER_ERRS 								   // 12
)

// 错误类型对应的报错信息
var ERROR_MSG = map[ERROR_TYPE]string{
	INVALID_S_NUM:              "%s: invalid start page %s. Start page should be lager than zero\n",
	INVALID_E_NUM:              "%s: invalid end page %s. End page should be lager than start page\n",
	INVALID_PAGE_LEN:           "%s: invalid page length %s\n",
	NOT_EXIST_INPUT_FILE:       "%s: input file \"%s\" does not exist\n",
	NOREADALE_INPUT_FILE:       "%s: input file \"%s\" exists but cannot be read\n",
	NOT_OPEN_INPUT_FILE:        "%s: could not open input file \"%s\"\n",
	NOT_START_PRINTER:          "%s: could not start a printer when run \"%s\"\n",
	NOT_OPEN_PIPE:              "%s: could not open pipe to \"%s\"\n",
	S_PAGE_GREATER_TOTAL_PAGES: "%s: start_page (%d) greater than total pages (%d) no output written\n",
	E_PAGE_GREATER_TOTAL_PAGES: "%s: end_page (%d) greater than total pages (%d), less output than expected\n",
	COMPLETE_PRINT_ERROR:       "%s: complete printing error with the printer %s\n",
	OTHER_ERRS:                 "%s: something wrong\n",
}

var progname string = EMPTY_STRING // 程序名

type Selpg_args struct { // 参数结构
	start_page  int
	end_page    int
	page_len    int
	paper_feed  bool
	print_dest  string
	in_filename string
}

func init_selpg(sp_args *Selpg_args) { // 初始化参数结构
	sp_args.start_page = -1
	sp_args.end_page = -1
	sp_args.page_len = 72
	sp_args.paper_feed = false
	sp_args.print_dest = EMPTY_STRING
	sp_args.in_filename = EMPTY_STRING
}

// --------------- 打印报错信息 ---------------
// 打印用法和报错信息并退出
func usage_and_exit(error_type ERROR_TYPE, message string) {
	flag.PrintDefaults()
	err_and_exit(error_type, message)
}

// 提供报错信息并退出
func err_and_exit(error_type ERROR_TYPE, message string) {
	fmt.Fprintf(os.Stderr, ERROR_MSG[error_type], progname, message)
	os.Exit(int(error_type))
}

// 页数出错
func page_ctr_err(error_type ERROR_TYPE, sa_page int, page_ctr int) {
	fmt.Fprintf(os.Stderr, ERROR_MSG[error_type], progname, sa_page, page_ctr)
}
// --------------------------------------------



// ------------- 判断输入参数是否有效 -----------------
// 判断数字是否在合理范围内 
func is_valid_number(number int) bool {
	return number > 0 && number < math.MaxInt32
}

// 判断start_page是否有效
func is_valid_start_page(start_page int) bool {
	return is_valid_number(start_page)
}

// 判断end_page是否有效
func is_valid_end_page(end_page int, start_page int) bool {
	return is_valid_number(end_page) && end_page >= start_page
}

// 判断是否为paper_feed类型
func is_paper_feed(page_feed bool) bool {
	return page_feed
}

// 判断page_len是否有效
func is_valid_page_len(page_len int) bool {
	return is_valid_number(page_len)
}

// 判断参数是否存在输入文件
func if_have_in_file(in_filename string) bool {
	return in_filename != EMPTY_STRING
}

// 检查start_page
func check_start_page(start_page int) {
	if !is_valid_start_page(start_page) {
		usage_and_exit(INVALID_S_NUM, strconv.Itoa(start_page))
	}
}

// 检查end_page
func check_end_page(end_page int, start_page int) {
	if !is_valid_end_page(end_page, start_page) {
		usage_and_exit(INVALID_E_NUM, strconv.Itoa(end_page))
	}
}

// 检查-f/-l参数
func check_paper_feed_and_page_len(paper_feed bool, page_len int) {
	if !is_paper_feed(paper_feed) && !is_valid_page_len(page_len) {
		usage_and_exit(INVALID_PAGE_LEN, strconv.Itoa(page_len))
	}
}

// 检查输入文件是否存在且可读
func check_in_file(in_filename string) {
	// 判断in_filename是否为空
	if if_have_in_file(in_filename) {
		// 文件不存在
		if _, err := os.Stat(in_filename); os.IsNotExist(err) {
			err_and_exit(NOT_EXIST_INPUT_FILE, in_filename)
		}
		// 文件不可读
		if _, err := ioutil.ReadFile(in_filename); err != nil {
			err_and_exit(NOREADALE_INPUT_FILE, in_filename)
		}
	}
}
// ---------------------------------------------

// 处理参数
func process_args(sp_args *Selpg_args) {
	// -------------- 判断参数是否有效 --------------
	// 检查 -s
	check_start_page(sp_args.start_page)
	// 检查 -e
	check_end_page(sp_args.end_page, sp_args.start_page)
	// 检查 -f / -l
	check_paper_feed_and_page_len(sp_args.paper_feed, sp_args.page_len)
	// 检查用户是否输入文件名且文件名是否有效
	check_in_file(sp_args.in_filename)
	// -------------------------------------
	// ONLY FOR DEBUDG
	// if !(sp_args.start_page > 0) {
	// 	fmt.Printf("start_page <= 0")
	// 	os.Exit(int(OTHER_ERRS))
	// }
	// if !(sp_args.end_page > 0) {
	// 	fmt.Printf("end_page <= 0")
	// 	os.Exit(int(OTHER_ERRS))
	// }
	// if !(sp_args.end_page >= sp_args.start_page) {
	// 	fmt.Printf("end_page < start_page")
	// 	os.Exit(int(OTHER_ERRS))
	// }
	// if !(sp_args.page_len >= 1) {
	// 	fmt.Printf("page_len <= 1")
	// 	os.Exit(int(OTHER_ERRS))
	// }
	// fmt.Printf("start_page = %d\n", sp_args.start_page)
	// fmt.Printf("end_page = %d\n", sp_args.end_page)
	// fmt.Printf("page_len = %d\n", sp_args.page_len)
	// fmt.Printf("paper_feed = %v\n", sp_args.paper_feed)
	// fmt.Printf("print_dest = %s\n", sp_args.print_dest)
	// fmt.Printf("in_filename = %s\n", sp_args.in_filename)

}

// 判断参数是否存在打印机名称
func if_have_print_dest(print_dest string) bool {
	return print_dest != EMPTY_STRING
}

// 打印EOF信息，不是错误
func EOF_msg() {
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)
}

func process_input(sp_args Selpg_args) {
	var fin_ptr *os.File // nil
	var fin *bufio.Reader
	var fout *bufio.Writer
	var stdinpipe io.WriteCloser
	var cmd *exec.Cmd

	// 从命令行或文件输入
	if if_have_in_file(sp_args.in_filename) { // 文件输入
		f, err := os.Open(sp_args.in_filename) // 打开文件
		if err != nil {
			err_and_exit(NOT_OPEN_INPUT_FILE, sp_args.in_filename)
		}
		fin_ptr = f
		fin = bufio.NewReader(f)
	} else { // 键盘输入
		fin = bufio.NewReader(os.Stdin)
	}

	// 输出到打印机或屏幕
	if if_have_print_dest(sp_args.print_dest) {
		var dest_flag string = fmt.Sprintf("-d%s", sp_args.print_dest)
		cmd = exec.Command("lp", dest_flag)
		stdin, err := cmd.StdinPipe() // 建立写入管道
		if err != nil {
			if fin_ptr != nil {
				fin_ptr.Close()
			}
			err_and_exit(NOT_OPEN_PIPE, fmt.Sprintf("lp %s", dest_flag))
		}
		stdinpipe = stdin
		fout = bufio.NewWriter(stdin)
		err = cmd.Start()

		if err != nil {
			if fin_ptr != nil {
				fin_ptr.Close()
			}
			stdin.Close()
			err_and_exit(NOT_START_PRINTER, fmt.Sprintf("lp %s", dest_flag))
		}
	} else {
		fout = bufio.NewWriter(os.Stdout)
	}

	var page_ctr int = 1 // 页数/开始打印第一页
	// 检测换页类型类型
	if !is_paper_feed(sp_args.paper_feed) {
		var line_ctr int = 0 // 行数
		for true {
			crc, err := fin.ReadString('\n') // 读取输入，知道遇到换行符
			if err == io.EOF {               // 处理err信息
				break
			} else if err != nil {
				panic(err)
			}
			line_ctr++                       // 行数增加
			if line_ctr > sp_args.page_len { // 超过页长度
				page_ctr++ // 增加一页
				line_ctr = 1
			}
			// 页数在指定范围内则输出
			if page_ctr >= sp_args.start_page && page_ctr <= sp_args.end_page {
				_, err := fout.Write([]byte(crc))
				if err != nil {
					panic(err)
				}
				fout.Flush()
			}
		}
	} else {
		for true {
			input_byte, err := fin.ReadByte()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			if input_byte == '\f' { // 换页走纸
				page_ctr++
				if page_ctr == sp_args.start_page {
					continue
				}
			}
			// 页数在指定范围内则输出
			if page_ctr >= sp_args.start_page && page_ctr <= sp_args.end_page {
				err := fout.WriteByte(input_byte)
				if err != nil {
					if fin_ptr != nil {
						fin_ptr.Close()
					}
					if if_have_print_dest(sp_args.print_dest) {
						stdinpipe.Close()
					}
					panic(err)
				}
				fout.Flush()
			}
		}
	}
	if page_ctr < sp_args.start_page {
		page_ctr_err(S_PAGE_GREATER_TOTAL_PAGES, sp_args.start_page, page_ctr)
	} else if page_ctr < sp_args.end_page {
		page_ctr_err(E_PAGE_GREATER_TOTAL_PAGES, sp_args.end_page, page_ctr)
	}
	// 正常EOF 没有出错
	fout.Flush()

	if if_have_print_dest(sp_args.print_dest) {
		stdinpipe.Close() // 关闭管道
		err := cmd.Wait()
		if err != nil {
			if fin_ptr != nil {
				fin_ptr.Close()
			}
			err_and_exit(COMPLETE_PRINT_ERROR, sp_args.print_dest)
		}
	}
	if fin_ptr != nil {
		fin_ptr.Close()
	}
	EOF_msg()

}

func main() {
	var sp_args Selpg_args
	init_selpg(&sp_args)
	progname = os.Args[0][strings.LastIndex(os.Args[0], string(os.PathSeparator))+1:]

	flag.IntVar(&sp_args.start_page, "s", -1, "begin printing from page \"start_page\"")
	flag.IntVar(&sp_args.end_page, "e", -1, "end printing after page \"end_page\"")
	flag.IntVar(&sp_args.page_len, "l", 72, "page_len: numbers of lines written on a piece of pa")
	flag.BoolVar(&sp_args.paper_feed, "f", false, "\n\tpage_type: if paper feed or not (default false)")
	flag.StringVar(&sp_args.print_dest, "d", EMPTY_STRING, "print_dest: name of printer")
	flag.Parse()

	if len(flag.Args()) > 0 {
		sp_args.in_filename = flag.Args()[0]
	}
	process_args(&sp_args)
	process_input(sp_args)
	os.Exit(0)
}

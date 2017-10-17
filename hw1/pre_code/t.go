package main

import (
	"bufio"
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
	EMPTY_STRING      = ""
	F            byte = 'f'
	L            byte = 'l'
)

// 错误类型的枚举
const (
	LACK_ARGS                  ERROR_TYPE = 1 + iota // 1
	LACK_S                                           // 2
	LACK_E                                           // 3
	INVALID_S_NUM                                    // 4
	INVALID_E_NUM                                    // 5
	INVALID_PAGE_LEN                                 // 6
	LACK_F                                           // 7
	LACK_DEST                                        // 8
	UNKNOWN_OPTION                                   // 9
	NOT_EXIST_INPUT_FILE                             // 10
	NOREADALE_INPUT_FILE                             // 11
	NOT_OPEN_INPUT_FILE                              // 12
	NOT_OPEN_PIPE                                    // 13
	S_PAGE_GREATER_TOTAL_PAGES                       // 15
	E_PAGE_GREATER_TOTAL_PAGES                       // 16
	OTHER_ERRS                                       // 17
)

// 错误类型对应的报错信息
var ERROR_MSG = map[ERROR_TYPE]string{
	LACK_ARGS:                  "%s: not enough arguments\n",
	LACK_S:                     "%s: 1st arg should be -sstart_page\n",
	INVALID_S_NUM:              "%s: invalid start page %s\n",
	LACK_E:                     "%s: 2nd arg should be -eend_page\n",
	INVALID_E_NUM:              "%s: invalid end page %s\n",
	INVALID_PAGE_LEN:           "%s: invalid page length %s\n",
	LACK_F:                     "%s: option should be \"-f\"\n",
	LACK_DEST:                  "%s: -d option requires a printer destination\n",
	UNKNOWN_OPTION:             "%s: unknown option %s\n",
	NOT_EXIST_INPUT_FILE:       "%s: input file \"%s\" does not exist\n",
	NOREADALE_INPUT_FILE:       "%s: input file \"%s\" exists but cannot be read\n",
	NOT_OPEN_INPUT_FILE:        "%s: could not open input file \"%s\"\n",
	NOT_OPEN_PIPE:              "%s: could not open pipe to \"%s\"\n",
	S_PAGE_GREATER_TOTAL_PAGES: "%s: start_page (%d) greater than total pages (%d) no output written\n",
	E_PAGE_GREATER_TOTAL_PAGES: "%s: end_page (%d) greater than total pages (%d), less output than expected\n",
}

var progname string = "" // 程序名

type Selpg_args struct { // 参数结构
	start_page  int
	end_page    int
	in_filename string
	page_len    int
	page_type   byte
	print_dest  string
}

func init_selpg(sp_args *Selpg_args) { // 初始化参数结构
	sp_args.start_page = -1
	sp_args.end_page = -1
	sp_args.in_filename = ""
	sp_args.page_len = 72
	sp_args.page_type = 'l'
	sp_args.print_dest = ""
}

// 提供报错信息用法并退出
func usage_and_exit(error_type ERROR_TYPE, str_msg string) {
	switch {
	case error_type >= LACK_ARGS && error_type <= LACK_E || error_type == LACK_F || error_type == LACK_DEST:
		fmt.Fprintf(os.Stderr, ERROR_MSG[error_type], progname)
	case error_type >= INVALID_S_NUM && error_type <= INVALID_PAGE_LEN || error_type == UNKNOWN_OPTION:
		fmt.Fprintf(os.Stderr, ERROR_MSG[error_type], progname, str_msg)
	default:
		fmt.Fprintf(os.Stderr, "For debug: error usage error type in usage_and_exit")
		os.Exit(int(OTHER_ERRS))
	}
	fmt.Fprintf(os.Stderr, "\nUSAGE: s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n")
	os.Exit(int(error_type))
}

// 与文件有关的报错
func file_err_exit(error_type ERROR_TYPE, file_name string) {
	fmt.Fprintf(os.Stderr, ERROR_MSG[error_type], progname, file_name)
	os.Exit(int(error_type))
}

// 页数出错
func page_ctr_err(error_type ERROR_TYPE, sa_page int, page_ctr int) {
	fmt.Fprintf(os.Stderr, ERROR_MSG[error_type], progname, sa_page, page_ctr)
}

// 判断字符串是否为数字且数字在合理范围内
func is_valid_str_number(str_num string, int_num *int) bool {
	if i, err := strconv.Atoi(str_num); err == nil && (i > 0 && i < math.MaxInt32) {
		*int_num = i
		return true
	}
	return false
}

// 处理参数
func process_args(ac int, av []string, sp_args *Selpg_args) {
	if ac < 3 {
		usage_and_exit(LACK_ARGS, "")
	}

	// ====================  -s  ===================
	if len(av[1]) < 2 || av[1][:2] != "-s" {
		usage_and_exit(LACK_S, "")
	}

	if !is_valid_str_number(av[1][2:], &sp_args.start_page) {
		usage_and_exit(INVALID_S_NUM, av[1][2:])
	}

	// ====================  -e  ===================
	if len(av[2]) < 2 || av[2][:2] != "-e" {
		usage_and_exit(LACK_E, "")
	}

	if !is_valid_str_number(av[2][2:], &sp_args.end_page) || sp_args.end_page < sp_args.start_page {
		usage_and_exit(INVALID_E_NUM, av[2][2:])
	}

	argno := 3
	// ================== 用户输入参数>3个，处理后面的参数  =============
	for ; argno < ac && av[argno][0] == '-'; argno++ {
		switch av[argno][1] {
		case 'l':
			// 数字长度不够且无效
			if len(av[argno]) < 2 || !is_valid_str_number(av[argno][2:], &sp_args.page_len) {
				usage_and_exit(INVALID_PAGE_LEN, av[argno][2:])
			}
		case 'f':
			if av[argno] == "-f" {
				sp_args.page_type = 'f'
			} else {
				usage_and_exit(LACK_F, "")
			}
		case 'd':
			// -d后有目标文件名
			if len(av[argno]) > 2 {
				sp_args.print_dest = av[argno][2:]
			} else {
				usage_and_exit(LACK_DEST, "")
			}
		default:
			usage_and_exit(UNKNOWN_OPTION, av[argno])
		}
	}
	// ==================== 处理非参数的输入，用户名 ===================
	if argno < ac { // 最多有一个
		sp_args.in_filename = av[argno]
		// 文件不存在
		if _, err := os.Stat(sp_args.in_filename); os.IsNotExist(err) {
			file_err_exit(NOT_EXIST_INPUT_FILE, sp_args.in_filename)
		}
		// 文件不可读
		if _, err := ioutil.ReadFile(sp_args.in_filename); err != nil {
			file_err_exit(NOREADALE_INPUT_FILE, sp_args.in_filename)
		}
	}
	// For debug
	// if !(sp_args.start_page > 0 && (sp_args.end_page > 0 && sp_args.end_page >= sp_args.start_page) &&
	//  sp_args.page_len > 1 && (sp_args.page_type == 'l' || sp_args.page_type == 'f')) {
	// 	 os.Exit(int(OTHER_ERRS))
	// }
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
	// if !(sp_args.page_type == 'l' || sp_args.page_type == 'f') {
	// 	fmt.Printf("sp_args.page_type = %c", sp_args.page_type)
	// 	os.Exit(int(OTHER_ERRS))
	// }
	// fmt.Printf("start_page = %d\n", sp_args.start_page)
	// fmt.Printf("end_page = %d\n", sp_args.end_page)
	// fmt.Printf("page_len = %d\n", sp_args.page_len)
	// fmt.Printf("page_type = %c\n", sp_args.page_type)
	// fmt.Printf("print_dest = %s\n", sp_args.print_dest)
	// fmt.Printf("in_filename = %s\n", sp_args.in_filename)

}

func process_input(sp_args Selpg_args) {
	var fin_ptr *os.File // nil
	var fin *bufio.Reader
	var fout *bufio.Writer
	var stdinpipe io.WriteCloser
	var cmd *exec.Cmd

	// 从命令行或文件输入
	if sp_args.in_filename != EMPTY_STRING { // 文件输入
		f, err := os.Open(sp_args.in_filename) // 打开文件
		if err != nil {
			file_err_exit(NOT_OPEN_INPUT_FILE, sp_args.in_filename)
		}
		fin_ptr = f
		fin = bufio.NewReader(f)
	} else { // 键盘输入
		fin = bufio.NewReader(os.Stdin)
	}

	// 输出到打印机或屏幕
	if sp_args.print_dest != EMPTY_STRING {
		var dest_flag string = fmt.Sprintf("-d%s", sp_args.print_dest)
		cmd = exec.Command("lp", dest_flag)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		stdinpipe = stdin
		fout = bufio.NewWriter(stdin)
		err = cmd.Start()

		if err != nil {
			file_err_exit(NOT_OPEN_PIPE, fmt.Sprintf("lp %s", dest_flag))
		}
	} else {
		fout = bufio.NewWriter(os.Stdout)
	}
	var page_ctr int = 1 // 页数/开始打印第一页
	// 检测页类型
	if sp_args.page_type == 'l' {
		var line_ctr int = 0 // 行数
		for true {
			// 读取输入，知道遇到换行符
			crc, err := fin.ReadString('\n')
			// 处理err信息
			if err == io.EOF {
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

	if sp_args.print_dest != EMPTY_STRING {
		stdinpipe.Close() // 关闭管道
		err := cmd.Wait()
		if err != nil {
			panic(err)
		}
	}
	if sp_args.in_filename != EMPTY_STRING {
		fin_ptr.Close()
	}
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)

}

func main() {
	var sp_args Selpg_args
	init_selpg(&sp_args)
	progname = os.Args[0][strings.LastIndex(os.Args[0], string(os.PathSeparator))+1:]

	process_args(len(os.Args), os.Args, &sp_args)
	process_input(sp_args)
	os.Exit(0)
}

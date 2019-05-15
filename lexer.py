# coding=utf-8

import sys
import string
import keywords

def pretreatment(file_name):
    '''
    预处理函数，将文件中的空格，换行等无关字符处理掉
    :param file_name: 处理文件的路径
    :return:
    '''
    try:
        fp_read = open(file_name, 'r')
        fp_write = open('file.tmp', 'w')
        sign = 0  # 是否遇到字符（非空格）
        # 不断读取，进行预处理
        while True:
            read = fp_read.readline()
            if not read:  # 读取结束
                break
            length = len(read)
            i = -1  # 索引
            while i < length - 1:
                i += 1
                if sign == 0:  # 还没有遇到字符
                    if read[i] == ' ':
                        continue
                if read[i] == '#':  # 后面均为注释，直接结束此行分析
                    break
                elif read[i] == ' ':  #
                    if sign == 1:
                        continue
                    else:
                        sign = 1
                        fp_write.write(' ')
                elif read[i] == '\t':  # table处理和空格相同
                    if sign == 1:
                        continue
                    else:
                        sign = 1
                        fp_write.write(' ')
                elif read[i] == '\n':  # 换行
                    if sign == 1:
                        continue
                    else:
                        fp_write.write(' ')  #
                        sign = 1
                elif read[i] == '"':
                    fp_write.write(read[i])
                    i += 1
                    while i < length and read[i] != '"':
                        fp_write.write(read[i])
                        i += 1
                    if i >= length:
                        break
                    fp_write.write(read[i])
                elif read[i] == "'":
                    fp_write.write(read[i])
                    i += 1
                    while i < length and read[i] != "'":
                        fp_write.write(read[i])
                        i += 1
                    if i >= length:
                        break
                    fp_write.write(read[i])
                else:
                    sign = 3
                    fp_write.write(read[i])
    except Exception:
        print(file_name, ': This FileName Not Found!')


def save(string):
    '''

    :param string:
    :return:
    '''
    if string in keywords.keywords.keys():
        if string not in keywords.signlist.keys():
            keywords.signlist[string] = keywords.keywords[string]
    else:
        try:
            float(string)
            save_const(string)
        except ValueError:
            save_var(string)


def save_var(string):
    '''

    :param string:
    :return:
    '''
    if string not in keywords.signlist.keys():
        if len(string.strip()) < 1:
            pass
        else:
            if is_signal(string) == 1:
                keywords.signlist[string] = 301
            else:
                keywords.signlist[string] = 501


def save_const(string):
    '''

    :param string:
    :return:
    '''
    if string not in keywords.signlist.keys():
        keywords.signlist[string] = 401


def save_error(string):
    '''

    :param string:
    :return:
    '''
    if string not in keywords.signlist.keys():
        keywords.signlist[string] = 501


def is_signal(s):
    '''

    :param s:
    :return:
    '''
    if s[0] == '_' or s[0] in string.ascii_letters:
        for i in s:
            if i in string.ascii_letters or i == '_' or i in string.digits:
                pass
            else:
                return 0
        return 1
    else:
        return 0


def recognition(filename):
    """

    :param filename:
    :return:
    """
    try:
        fp_read = open(filename, 'r')
        string = ""
        sign = 0
        while True:
            read = fp_read.read(1)
            if not read:
                break

            if read == ' ':
                if len(string.strip()) < 1:
                    sign = 0
                    pass
                else:
                    if sign == 1 or sign == 2:
                        string += read
                    else:
                        save(string)
                        string = ""
                        sign = 0
            elif read == '(':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save('(')
            elif read == ')':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save(')')
            elif read == '[':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save('[')
            elif read == ']':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save(']')
            elif read == '{':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save('{')
            elif read == '}':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save('}')
            elif read == '<':
                save(string)
                string = ""
                save('<')
            elif read == '>':
                save(string)
                string = ""
                save('>')
            elif read == ',':
                save(string)
                string = ""
                save(',')
            elif read == "'":
                string += read
                if sign == 1:
                    sign = 0
                    save_const(string)
                    string = ""
                else:
                    if sign != 2:
                        sign = 1
            elif read == '"':
                string += read
                if sign == 2:
                    sign = 0
                    save_const(string)
                    string = ""
                else:
                    if sign != 1:
                        sign = 2
            elif read == ':':
                if sign == 1 or sign == 2:
                    string += read
                else:
                    save(string)
                    string = ""
                    save(':')
            elif read == '+':
                save(string)
                string = ""
                save('+')
            elif read == '=':
                save(string)
                string = ""
                save('=')
            else:
                string += read

    except Exception as e:
        print(e)


def main():
    file_name = "C:\\Data\\Projects\\PycharmProjects\\Network\\test\\test.txt"
    pretreatment(file_name)
    recognition('file.tmp')
    for i in keywords.signlist.keys():
        print("(", keywords.signlist[i], ",", i, ")")


if __name__ == '__main__':
    main()

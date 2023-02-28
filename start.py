# ! /usr/bin/env python3
# -*- coding: utf-8 -*-
# @author  : BeYoung
# @time    : 2023/2/24

import argparse
import multiprocessing
import os
import pathlib
import subprocess
import time

PATH = []
TASKS = []


def parser_arg():
    parser = argparse.ArgumentParser()
    parser.add_argument("--kill",
                        nargs="?",
                        type=int,
                        default=0,
                        help="kill all service")
    parser.add_argument("--run",
                        nargs="?",
                        type=int,
                        default=0,
                        help="run all service, but not execute command: make dep, make init")
    parser.add_argument("--main",
                        nargs="?",
                        type=int,
                        default=0,
                        help="run all service, and execute command: make dep, make init")
    args = parser.parse_args()
    if args.kill != 0:
        kill_all_main()
        os._exit(0)
    elif args.run != 0:
        get_all_makefile()
        make_all_run()
        time.sleep(3)
        wait_all_main()
    elif args.main != 0:
        get_all_makefile()
        make_dep()
        make_init()
        make_all_run()
        time.sleep(3)
        wait_all_main()


def on_exit(signum, frame):
    for task in TASKS:
        task: multiprocessing.Process
        task.close()


def kill_all_main():
    # 因为要获取 shell 的输出，所以使用 subprocess.getoutput 执行 shell 命令
    pids = subprocess.getoutput("pidof main").split(" ")
    for pid in pids:
        os.system(f"kill {pid}")


def run_main(path):
    path: pathlib.Path
    os.chdir(path.parent)
    print(f"Now path: {os.getcwd()}")
    print(f"make run: {path.cwd()}")
    os.system(f"make run >> ~/log/{path.name}.log")


def make_all_run():
    for path in PATH:
        task = multiprocessing.Process(target=run_main, args=(path,))
        TASKS.append(task)
        task.daemon = True
        task.start()


def wait_all_main():
    for task in TASKS:
        task: multiprocessing.Process
        task.join()


def get_all_makefile():
    global PATH
    cwd = pathlib.Path.cwd()
    for path in cwd.iterdir():
        if path.is_dir() and not path.is_file():
            PATH += [makefile for makefile in path.iterdir() if makefile.name == "Makefile" and makefile.is_file()]


def make_dep():
    for makefile in PATH:
        makefile: pathlib.Path
        os.chdir(makefile.parent)
        print(f"Now path: {os.getcwd()}")
        print(f"make dep: {makefile.cwd()}")
        os.system(f"make dep >> ~/log/{makefile.parent.name}.log")


def make_init():
    for makefile in PATH:
        makefile: pathlib.Path
        os.chdir(makefile.parent)
        print(f"Now path: {os.getcwd()}")
        print(f"make init: {makefile.cwd()}")
        os.system(f"make init >> ~/log/{makefile.parent.name}.log")


def make_install():
    for makefile in PATH:
        makefile: pathlib.Path
        os.chdir(makefile)
        print(f"Now path: {os.getcwd()}")
        print(f"make install: {makefile.cwd()}")
        os.system(f"make install >> ~/log/{makefile.parent.name}.log")


def main():
    parser_arg()
    get_all_makefile()
    make_dep()
    make_init()
    make_all_run()
    time.sleep(3)
    wait_all_main()


if __name__ == '__main__':
    main()

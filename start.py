import argparse
import multiprocessing
import os
import pathlib
import signal
import subprocess
import time

DIRS = []
TASKS = []


def parser_arg():
    parser = argparse.ArgumentParser()
    parser.add_argument("--kill",
                        nargs="?",
                        type=int,
                        default=0,
                        help="kill all service")
    args = parser.parse_args()
    if args.kill != 0:
        kill_all_main()
        os._exit(0)


def on_exit(signum, frame):
    for task in TASKS:
        task: multiprocessing.Process
        task.close()


def kill_all_main():
    pids = subprocess.getoutput("pidof main").strip(" ")
    for pid in pids:
        os.kill(int(pid), signal.SIGKILL)


def get_all_main():
    cwd = pathlib.Path.cwd()
    DIRS.append(cwd.joinpath("network"))
    print(f'{cwd.joinpath("network")}')
    path = cwd.joinpath("apps")
    for p in path.iterdir():
        print(f"{p}")
        DIRS.append(p)


def run_main(path):
    path: pathlib.Path
    os.chdir(path)
    print(f"Now path: {os.getcwd()}\n"
          f"RUN {path.name}, PID: {os.getpid()}")
    os.system(f"go run {path}/main.go >> ~/log/{path.name}.log")


def run_all_main():
    for path in DIRS:
        task = multiprocessing.Process(target=run_main, args=(path,))
        TASKS.append(task)
        task.daemon = True
        task.start()


def wait_all_main():
    for task in TASKS:
        task: multiprocessing.Process
        task.join()


def main():
    path = pathlib.Path.cwd()
    print(f"Now path: {os.getcwd()}\n"
          f"RUN {path.name}, PID: {os.getpid()}")

    parser_arg()
    get_all_main()
    run_all_main()
    time.sleep(3)
    wait_all_main()


if __name__ == '__main__':
    main()

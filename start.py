import argparse
import multiprocessing
import os
import pathlib
import time

DIRS = []
TASKS = []
PID = []


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
    with open(f"{pathlib.Path().home()}/log/pid.txt", encoding="utf8") as file:
        for pid in file.readlines():
            print(f"closing: {pid}")
            os.system(f"kill -9 {pid}")
    os.system(f"echo ' ' > {pathlib.Path().home()}/log/pid.txt")


def save_all_pid():
    try:
        with open(f"{pathlib.Path().home()}/log/pid.txt", mode="a", encoding="utf8") as file:
            for p in PID:
                print(f"saving: {p}")
                file.write(f"{p} /n")
    except Exception as e:
        print(e)


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
    PID.append(os.getpid())
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
    parser_arg()
    get_all_main()
    run_all_main()
    time.sleep(3)
    save_all_pid()
    wait_all_main()


if __name__ == '__main__':
    main()

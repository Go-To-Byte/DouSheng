import argparse
import multiprocessing
import os
import pathlib
import time

DIRS = []
TASKS = []


def on_exit(signum, frame):
    for task in TASKS:
        task: multiprocessing.Process
        task.close()


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
    print(f"Now path: {os.getcwd()}")
    print(f"go run {path}/main.go")
    os.system(f"go run {path}/main.go >> ~/log/{path.name}.log")


def run_all_main():
    for path in DIRS:
        task = multiprocessing.Process(target=run_main, args=(path,))
        TASKS.append(task)
        task.daemon = True
        task.start()
        print(f"Run: {path}: {task.pid}")


def wait_all_main():
    for task in TASKS:
        task: multiprocessing.Process
        task.join()


def main():
    parser = argparse.ArgumentParser()

    get_all_main()
    run_all_main()
    time.sleep(3)
    wait_all_main()


if __name__ == '__main__':
    main()

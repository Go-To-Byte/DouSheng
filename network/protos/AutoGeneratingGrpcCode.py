import logging
import os
import pathlib

CMD = "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "


def get_working_path() -> str:
    logging.info(f"work file path is: {__file__}")
    return __file__


def generate_code():
    base_path = get_working_path()
    path = pathlib.Path(base_path).parent
    logging.info(f"work dir path is: {path}")
    for file in path.iterdir():
        if file.suffix == ".py" or file.suffix == ".go":
            continue
        else:
            filename = file.name
            logging.info(f"find file{file}\ngenerating code: {filename}")
            os.system(f"{CMD} {filename}")


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    logging.info("Generating grpc code...")
    generate_code()

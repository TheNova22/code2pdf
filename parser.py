import argparse


class Parser:

    def __init__(self):
        self.parser = argparse.ArgumentParser(
                description=(
                    "Convert any given code file to pdf and save it"),
                epilog="Author:jayantsogikar@gmail.com"
            )
        self.parser.add_argument(
            "--f",
            help="list of files that are required to converted",
            nargs="*",
            type=str)
        self.parser.add_argument(
            "--u",
            help="list of raw github urls that are required to converted",
            nargs="*",
            type=str)
        self.parser.add_argument(
            "--a",
            help="Author's name tag. It adds it at the start of the file as a comment",
            default="",
            type=str)
        self.parser.add_argument(
            "--n",
            help="Name of the pdf to be saved as",
            default="",
            type=str)
        self.parser.add_argument(
            "--s",
            help="the style name for highlighting. Default is set as zenburn",
            type=str,
            default="zenburn")

    def get_args(self):
        return self.parser.parse_args()
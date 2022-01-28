# -*- coding: UTF-8 -*

from PIL import Image
import argparse
import os
import sys

from PIL import Image

bounding_box1 = (0, 8, 50, 25)
bounding_box2 = (50, 8, 100, 25)

def rename(input_path):
    output_path = input_path.replace("_out", "")
    with Image.open(input_path) as image:
        os.remove(input_path)
        image.save(output_path)


def crop(input_path):
    items = input_path.rpartition("/")
    with Image.open(input_path) as image:
        crop1 = image.crop(bounding_box1)
        crop2 = image.crop(bounding_box2)
        crop1.save(items[0] + "/" + items[2].replace(".png", "1.png"))
        crop2.save(items[0] + "/" + items[2].replace(".png", "2.png"))


def resize(input_path, output_path, scale):

    file_name = input_path.rpartition("/")[2]
    print(input_path.rpartition("/"))

    with Image.open(input_path) as image:
        # image.show()
        ori_size = image.size  # this seems a tuple.
        image = image.resize((int(ori_size[0] * scale), int(image.size[1] * scale)), Image.ANTIALIAS)
        outfile = output_path + "/" + file_name
        print(outfile)
        image.save(outfile)


def find_files(input_path):
    img_file_list = []
    for name in os.listdir(input_path):
        file_abs_path = input_path + "/" + name
        # print(file_full_path)
        if os.path.isfile(file_abs_path):
            if file_abs_path.endswith('.png') or file_abs_path.endswith('.jpg'):
                img_file_list.append(file_abs_path)
        elif os.path.isdir(file_abs_path):
            img_file_list.extend(find_files(file_abs_path))
    return img_file_list


if __name__ == '__main__':

    parser = argparse.ArgumentParser(usage='please use plistCutter.py -h to get help information.')
    parser.add_argument('-input_path', help='Specify a directory to get imgs.', default=".")
    parser.add_argument('-output_path', help='Specify a directory to output file.', default=".")
    parser.add_argument('-scale', help='Specify a the scale ratio of the scale operation.', default="1")
    parser.add_argument('-option', help='Specify which operation to operate.', default="rename")

    args = parser.parse_args()
    input_path = args.input_path
    output_path = args.output_path
    scale = float(args.scale)
    option = args.option

    print("the scale is ", args.scale)

    file_list = find_files(input_path)
    if option == "resize":
        if input_path is None and not os.path.isdir(input_path):
            print('error: %s is not an valid dir or doesn\'t exist.')
            sys.exit(0)
        # print("len: ", str(len(img_file_path_list)))
        for file in file_list:
            if os.path.exists(file):
                resize(file, output_path, scale)
        print("Analyze...Done.")
    elif option == "rename":
        for file in file_list:
            rename(file)
    elif option == "crop":
        for file in file_list:
            crop(file)


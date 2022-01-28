# -*- coding: UTF-8 -*

from xml.etree import ElementTree
from PIL import Image
import argparse
import os
import sys


class PlistParser:
    def __init__(self, plist_path, image_path, output_dir_path="./"):
        self.plist_path = plist_path
        self.image_path = image_path
        self.output_dir_path = output_dir_path

    def convert_tree_to_dict(self, tree):
        d = {}  # 返回值是一个dict
        for index, item in enumerate(tree):  # 遍历整棵XML树
            if item.tag == 'key':  # 如果该item的tag为'key'
                # 根据下一个结点的tag值不同，放在dict的不同位置上
                if tree[index + 1].tag == 'string':
                    d[item.text] = tree[index + 1].text
                elif tree[index + 1].tag == 'true':
                    d[item.text] = True
                elif tree[index + 1].tag == 'false':
                    d[item.text] = False
                elif tree[index + 1].tag == 'dict':
                    d[item.text] = self.convert_tree_to_dict(tree[index + 1])  # 递归下去
            elif item.tag == 'dict' and item.getchildren()[0].text == 'frames':
                d = self.convert_tree_to_dict(item)

        return d

    def split_image_from_plist(self):
        output_dir = self.output_dir_path
        if output_dir is None:
            output_dir = self.plist_path.replace('.plist', '')
            if not os.path.isdir(output_dir):  # 如果不存在该目录
                os.mkdir(output_dir)  # 新建一个目录

        src_img = Image.open(self.image_path)  # 打开图像
        plist_content_str = open(self.plist_path, 'r').read()  # 读取plist文件
        plist_xml_tree = ElementTree.fromstring(plist_content_str)  # 转换成XML树
        plist_dict = self.convert_tree_to_dict(plist_xml_tree)  # 获取dict

        for key, value in plist_dict['frames'].items():
            pos_str = str(value['frame'])  # 获取表示位置的str
            pos_rect = pos_str.replace("{", "").replace("}", "").split(",")
            source_color_rect = str(value["sourceColorRect"]).replace("{", "").replace("}", "").split(",")
            source_size = str(value["sourceSize"]).replace("{", "").replace("}", "").split(",")
            source_offset = str(value["offset"]).replace("{", "").replace("}", "").split(",")

            width = int(pos_rect[3] if value['rotated'] else pos_rect[2])
            height = int(pos_rect[2] if value['rotated'] else pos_rect[3])
            bounding_box = (
                int(pos_rect[0]),
                int(pos_rect[1]),
                int(pos_rect[0]) + width,
                int(pos_rect[1]) + height
            )
            cropped_img = src_img.crop(bounding_box)
            # 这里不能简单用rotate
            # 简单用transpose
            if value['rotated']:
                cropped_img = cropped_img.transpose(Image.ROTATE_90)
            background_img = Image.new('RGBA', (int(source_size[0]), int(source_size[1])))
            background_img.paste(cropped_img, (int(source_color_rect[0]), int(source_color_rect[1])), mask=None)

            outfile = output_dir + "/" + key
            background_img.save(outfile)



# 在该路径中寻找所有plist文件
def find_plist_file(dir_path):

    plist_file_list = []
    for short_name in os.listdir(dir_path):
        file_full_path = dir_path + "/" + short_name
        # print(file_full_path)
        if os.path.isfile(file_full_path):
            if file_full_path.endswith('.plist') and not file_full_path.startswith('.'):
                plist_file_list.append(file_full_path)
        elif os.path.isdir(file_full_path):
            plist_file_list.extend(find_plist_file(file_full_path))

    return plist_file_list


if __name__ == '__main__':
    parser = argparse.ArgumentParser(usage='please use plistCutter.py -h to get help information.')
    parser.add_argument('-dir', help='Specify a directory file path you would like to find')


    args = parser.parse_args()
    dir = args.dir

    if dir is None and not os.path.isdir(dir):
        print('error: %s is not an valid dir or doesn\'t exist.')
        sys.exit(0)

    plist_file_list = find_plist_file(dir)
    print("total plist file num: " + str(len(plist_file_list)))
    for plist_file in plist_file_list:
        plist_path = plist_file
        png_path = plist_file.replace('.plist', '.png')
        if os.path.exists(png_path):
            output_dir = plist_file.rpartition("/")[0] + "/res"
            if not os.path.exists(output_dir):
                os.mkdir(output_dir)
            print(output_dir + "\n" + plist_file)
            PlistParser(plist_path, png_path, output_dir).split_image_from_plist()

    print("Analyze...Done.")

import hashlib
import socket
import json
import sys
from collections import defaultdict

HOST = 'localhost'
PORT = 8888

conn: socket


def send_command_to_server(cmd):
    cmd += "\n"
    bytes_str = str.encode(cmd)
    try:
        conn.send(bytes_str)
        data = conn.recv(4096)
    except:
        print("\nCan't connect to the server. Please, try again later")
        conn.close()
        sys.exit()
    return data.decode('utf-8').strip("\n")


def try_login():
    print("Input login: ", end='')
    login = str(input())
    if login == "admin":
        return False, login

    print("Input password: ", end='')
    password = hashlib.sha256(str(input()).encode()).hexdigest()

    response = send_command_to_server("verify " + login + " " + password)
    return response == "success", login


def character_menu():
    pass


def character_view(info):
    print("\nCharacter view:")
    print("1. Info")
    print("2. Inventory")
    print("3. Exit")
    print("Input: ", end='')
    command = input()
    print()
    if command == "1":
        print("Character Info:")
        print("    Race: " + str(info[0][0]))
        print("    Level: " + str(info[0][1]))
        print("\nPress Enter to continue")
        input()
    elif command == "2":
        print("Inventory:")
        step = 0
        for item in info:
            step += 1
            print("    " + str(step) + ". " + str(item[2]))
        print("\nPress Enter to continue")
        input()
    elif command == "3":
        return False
    else:
        print("Wrong input\n")
    return True


def characters_view(character_dict):
    step = 0
    print("\nChoose your character:")
    for character in character_dict:
        step += 1
        print(str(step) + ". " + character)
    print("0. Exit")
    print("Input: ", end='')
    character_id = input()
    if not character_id.isdigit():
        print("Wrong input\n")
    elif 0 < int(character_id) <= len(character_dict):
        step = 0
        for key in character_dict:
            step += 1
            if str(step) == character_id:
                while character_view(character_dict[key]):
                    pass
    elif character_id == "0":
        return False
    return True


def get_characters(login):
    print("\nSuccessful login!")
    response = send_command_to_server("get_characters " + login)
    characters = json.loads(response)
    characters_dict = defaultdict(list[tuple])
    for character in characters:
        characters_dict[character["nickname"]].append((character["race"], character["level"], character["itemName"]))

    while characters_view(characters_dict):
        pass


def auth():
    print("Authentication menu:")
    print("1. Try login")
    print("2. Main menu")
    print("Input: ", end='')
    command = input()
    print()
    if command == "1":
        login_response = try_login()
        if login_response[0]:
            try:
                get_characters(login_response[1])
            except:
                print("\nBroken account\n")
        elif login_response[1] == "admin":
            print("\nNo admin panel here\n")
        else:
            print("\nWrong login or password\n")
    elif command == "2":
        return False
    else:
        print("Wrong input\n")
    return True


def menu():
    print("Menu:")
    print("1. Authentication")
    print("2. Quit")
    print("Input: ", end='')
    command = input()
    print()
    if command == "1":
        while auth():
            pass
    elif command == "2":
        return False
    else:
        print("Wrong input\n")
    return True


conn = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
try:
    conn.connect((HOST, PORT))
except:
    print("\nCan't connect to the server. Please, try again later")
    conn.close()
    sys.exit()

while menu():
    pass

send_command_to_server("quit")
conn.close()

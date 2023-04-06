#!/usr/bin/python3
import requests
import socket

URL = input("URL with http(s): ")
urlList = []
isFollowed = {}

hostname = URL.split("//")[-1].split("/")[0]
ip_address = socket.gethostbyname(hostname)
print(f"This is the IP address for the URL: {ip_address}")

def checkUrlList(URL):
    if URL in urlList:
        return True
    else:
        return False

def isFollowedCheck(URL):
    for entry in isFollowed.keys():
        if URL != entry:
            return False
        else:
            if isFollowed[URL] == "yes":
                return True
            else:
                return False
urlList.append(URL)

for URL in urlList:
    if isFollowedCheck(URL) != True:
        try:
            page = requests.get(URL, timeout=10)
        except Exception as e:
            print(f"An error occurred while accessing {URL}")
            continue

        isFollowed[URL] = "yes"

        start = "http"
        for line in page.text.split("\n"):
            if "http" in line:
                if hostname in line:
                    try:
                        end = line.index('"', line.index(start) + len(start))
                        sliced = line[line.index(start):end]
                    except ValueError:
                        try:
                            end = line.index("'", line.index(start) + len(start))
                            sliced = line[line.index(start):end]
                        except ValueError:
                            continue
                    parsedURL = sliced
                    if checkUrlList(parsedURL) == False:
                        urlList.append(parsedURL)
                        isFollowed[parsedURL] = "no"
                        print(parsedURL)

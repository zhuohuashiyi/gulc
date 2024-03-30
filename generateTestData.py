import random
import json
import pyperclip



# generateUniqueArray 生成不重复数组        
def generateUniqueArray(length, lower, upper):
    res = []
    i = 0
    while i < length:
        a = random.randint(lower, upper)
        if a not in res:
            res.append(a)
            i += 1
    pyperclip.copy(json.dumps(res))
    
    
def generateDemo():
    # skills = []
    # mustExisted = set([])
    # for i in range(16):
    #     skills.append(chr(65 + i))
    #     mustExisted.add(skills[i])
    # peoples = []
    # for i in range(59):
    #     people = []
    #     n = random.randint(1, 16)
    #     j = 0
    #     while j < n:
    #         skill = skills[random.randint(0, 15)]
    #         if skill not in people:
    #             if skill in mustExisted:
    #                 mustExisted.remove(skill)
    #             people.append(skill)
    #             j += 1
    #     peoples.append(people)
    # if len(mustExisted) != 0:
    #     peoples.append(list(mustExisted))
    # pyperclip.copy((json.dumps(skills) + '\n' + json.dumps(peoples)).replace('[', '{').replace(']', '}'))
    # str = ""
    # for i in range(10 ** 2):
    #     a = random.randint(0, 25)
    #     str += chr(97 + a)
    # pyperclip.copy(str)
    # res = []
    # for i in range(50):
    #     arr = []
    #     for j in range(50):
    #         arr.append(random.randint(-1, 1))
    #     res.append(arr)
    # res[0][0] = res[-1][-1] = 0
    # pyperclip.copy(json.dumps(res))
    res = ""
    for i in range(15):
        res += str(random.randint(0, 1))
    pyperclip.copy(json.dumps(res))
    
    
if __name__ == "__main__":
    #generateUniqueArray(10 ** 4, 1, 10 ** 9)
    generateDemo()
    
    
    

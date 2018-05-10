import requests


def init_user():
    data = {
        'firstname': 'zhyq',
        'lastname': 'san',
        'email': 'zhyq@gmail.com',
        'password': '1123456',
    }
    requests.post('http://localhost:8000/users', json={'data': data})


AUTH_HEADER = {
    'Authorization': u'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InpoeXFAZ21haWwuY29tIiwicm9sZSI6Im1lbWJlciIsImV4cCI6MTUyNTkzMjcwNiwiaXNzIjoiYWRtaW4ifQ.fX1nQCTNM_U2jkSsFffkpghCgqvGrvrYaLkMOeU9gXx4NHJARCbiAdguJ5wYfOgd8BIUJX-ATmGggbLq4YIFvh9VUd6lsOZatFskynQ3XKXtJBEHTxHrpFFlexBx35PIA713x6R4_Wcilbs5wwQ4C8mXBPNMT3ahh65797d3k9s'
}


def login_user():
    data = {
        'email': 'zhyq@gmail.com',
        'password': '1123456',
    }
    resp = requests.post(
        'http://localhost:8000/users/login', json={'data': data})
    print resp.json()['data']['token']


def list_bookmarks():
    resp = requests.get('http://localhost:8000/bookmarks', headers=AUTH_HEADER)
    print resp.json()


if __name__ == '__main__':
    list_bookmarks()
    # login_user()

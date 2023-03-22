import requests, os, pytest

def pytest_configure():
    pytest.TOKEN = ""
    pytest.first_flight_id = ""
    
def test_init():
    req = requests.get("http://localhost:8080/init-flight")
    assert req.status_code == 200 

# def test_login():
#     req = requests.post(url="http://localhost:8080/login", json={"username": "naz1", "password": "123"})
#     pytest.TOKEN = req.headers['Bearer']
#     assert pytest.TOKEN.startswith("ey")

def test_create_flight():
    req = requests.post(url="http://localhost:8080/flight", json = {
        "date": "2023-05-06T12:00:42.123Z",
        "endPlace": "London",
        "startPlace": "Budapest",
        "capacity": 200,
        "price": 150,
        "freeSeats": 58
    })
    assert req.status_code == 201

def test_get_all_flights():
    req = requests.get(url="http://localhost:8080/flight")
    pytest.first_flight_id = req.json()[0]['id']
    #checking only username and passoword, not ID
    assert list(map(lambda x: {x["date"], x["endPlace"], x["startPlace"],  x["capacity"],  x["price"],  x["freeSeats"]},req.json())) == [{
        "2023-05-06T12:00:42.123Z",
        "London",
        "Budapest",
        200,
        150,
        58
    }]

def test_get_flight_by_id():
    req = requests.get(url="http://localhost:8080/get-flight-by-id?id="+pytest.first_flight_id)
    assert req.json()[0] == {
        'id': pytest.first_flight_id,
        'capacity': 200,
        'date': '2023-05-06T12:00:42.123Z',
        'endPlace': 'London',
        'freeSeats': 58,
        'price': 150, 
        'startPlace': 'Budapest'}

# def test_delete_flight():
#     req = requests.delete(url="http://localhost:8080/flight/"+pytest.first_flight_id)

def test_update_flight():
    req = requests.patch(url="http://localhost:8080/flight/"+pytest.first_flight_id, json={
        'id': pytest.first_flight_id,
        'capacity': 200,
        'date': '2023-12-31T12:00:42.123Z',
        'endPlace': 'London',
        'freeSeats': 50,
        'price': 150, 
        'startPlace': 'Budapest'})
    req = requests.get(url="http://localhost:8080/get-flight-by-id?id="+pytest.first_flight_id)
    assert req.json()[0] == {
        'id': pytest.first_flight_id,
        'capacity': 200,
        'date': '2023-12-31T12:00:42.123Z',
        'endPlace': 'London',
        'freeSeats': 50,
        'price': 150, 
        'startPlace': 'Budapest'}
    

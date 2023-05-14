import axios from 'axios';

export default class APIService{
    static URL = "http://127.0.0.1:8000";

    static async Register(username, password, email, firstName, lastName, phone, roleNum, address, city, country) {
      try {
        const response = await axios.post(APIService.URL + '/user', {
          username: username,
          name: firstName,
          surname:lastName,
          phoneNumber: phone,
          email: email,
          password: password,
          role: roleNum,
          address: address,
          city: city,
          country: country

        }, {
          headers: {
            'Content-Type': 'application/json'
          }
        });
        return response.data;
      } catch (error) {
        throw error
      }
    }

    static async Login(username, password) {
        try {
          const response = await axios.post(APIService.URL + '/login', {
            username: username,
            password: password,
          }, {
            headers: {
              'Content-Type': 'application/json'
            }
          });
          return response.data;
        } catch (error) {
          throw error
        }
      }
}
  
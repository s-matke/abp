import axios from 'axios';

export default class APIService{
    static URL = "http://127.0.0.1:8000/";

    static async Register(username, password, email, firstName, lastName, phone) {
      try {
        const response = await axios.post(APIService.URL + 'api/user/', {
          username: username,
          name: firstName,
          surname:lastName,
          phoneNumber: phone,
          email: email,
          password: password,
          role: "user",
          address: "123 Main St",
          city: "Anytown",
          country: "USA"

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
  
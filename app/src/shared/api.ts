import axios, { AxiosInstance } from "axios";

class Api {
  private static axiosInstance: AxiosInstance;
  static init() {
    this.axiosInstance = axios.create({
      // Endpoint URL
      baseURL: "http://localhost:5000",
    });
  }
}

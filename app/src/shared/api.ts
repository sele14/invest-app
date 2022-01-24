import axios, { AxiosInstance } from "axios";

interface ResponseType {
  hey: string;
}

class Api {
  private static axiosInstance: AxiosInstance;

  static init() {
    this.axiosInstance = axios.create({
      // Endpoint URL
      baseURL: "http://localhost:5000",
    });
  }

  static async get<ResponseType>(url: string) {
    return await Api.axiosInstance.get(url);
  }
  static async post<ResponseType, DataType>(url: string, data?: DataType) {}
}

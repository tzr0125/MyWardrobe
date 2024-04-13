'use server';
import Axios, { AxiosResponse } from 'axios';
import { User } from 'next-auth';

export async function authorize(credentials: Partial<Record<"email" | "password", unknown>>, request: Request) {
    Axios.defaults.baseURL='http://localhost:8080/';
    Axios.defaults.withCredentials = true;
    console.log(Axios.defaults.baseURL);
    const response = await Axios(
        {
            url: '/api/test/getauth',
            method: 'POST'
        }
    ).then(response => {
        setAuthHeader(response);
        return response;
    })
    const user: User = {
        "id": "111",
    }
    if (process.browser) {
        console.log('Code is running on the client');
      } else {
        console.log('Code is running on the server');
      }
    if (response !== undefined) {
        console.log("has return user");
        return user;
    }
    return null;
}

function setAuthHeader(response: AxiosResponse) {
    const key = response.data.key;
    console.log("data", response.data);
    if (key) {
        Axios.defaults.headers.common['Authorization'] = `Bearer ${key}`;
        console.log("aaa", Axios.defaults.headers.common['Authorization']);

        Axios.defaults.headers.common['X-csrf-token'] = response.headers['set-cookie']
    } else {
        delete Axios.defaults.headers.common['Authorization'];
    }
}

export async function authorized({ auth, request: { nextUrl } }) {

    if (process.browser) {
        console.log('Code is running on the client');
      } else {
        console.log('Code is running on the server');
      }

    const isLoggedIn = !!auth?.user;
    console.log("user", auth?.user);
    console.log("baseURL", Axios.defaults.baseURL)
    const isOnDashboard = nextUrl.pathname.startsWith('/index');
    if (isOnDashboard) {
      if (isLoggedIn) return true;
      return false; // Redirect unauthenticated users to login page
    } else if (isLoggedIn) {
      return Response.redirect(new URL('/index', nextUrl));
    }
    return true;
}
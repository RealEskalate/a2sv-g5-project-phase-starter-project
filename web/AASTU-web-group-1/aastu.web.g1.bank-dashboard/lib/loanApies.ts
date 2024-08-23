import { getSession } from "next-auth/react";

const baseUrl = `https://bank-dashboard-o9tl.onrender.com`

export  async function getLoansAll(page:number, size:number) {
    try {
        const session = await getSession();
        const accessToken = session?.user.accessToken;
        const response = await fetch(`${baseUrl}/active-loans/all?page=${page}&size=${size}`, {
            method: "GET",
            cache: "reload",
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
            body: null,
        }).then((res)=>(res.json()));

        if (response.success){
            return response.data;
        }else{
            throw new Error("failed to get data");
        }
    } catch (error) {
      console.error("An error occurred on card:", error);
    }
}

export async function getDetailData() {
    try{
        const session = await getSession();
        const accessToken = session?.user.accessToken;
        const response = await fetch(`${baseUrl}/active-loans/detail-data`, {
            method: "GET",
            cache:"no-cache",
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
            body: null
        }).then((res)=>(res.json()));

        if(response.success){
            console.log(response.data);
            return response.data;
        }else{
            throw new Error(`failed to fetch the data: ${response}`)
        }
    }catch (error){
        console.error(error);
    }
}
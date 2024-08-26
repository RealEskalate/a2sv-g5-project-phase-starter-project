import axios from "axios";
const baseUrl = "https://a2svwallets.onrender.com";
const activeloansall = async (token: string, size: number, page: number) => {
  const response = await axios.get(
    baseUrl + `/active-loans/all?page=${page}&size=${size}`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
  return response.data.data;
};

const activeloansdetaildata = async (token: string) => {
  const response = await axios.get(baseUrl + "/active-loans/detail-data", {
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  return response.data.data;
};

const activeloansmyloans = async (token: string) => {
  const response = await axios.get(baseUrl + "/active-loans/my-loans", {
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  return response.data.data;
};

const activeloansid = async (loanid: string, token: string) => {
  const response = await axios.get(baseUrl + `/active-loans/${loanid}`, {
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
  return response.data.data;
};

const activeloansidapprove = async (loanid: string, token: string) => {
  const respons7 = await axios.post(
    baseUrl + `/active-loans/${loanid}/approve`,
    {},
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
};

const activeloansidreject = async (loanid: string, token: string) => {
  const response = await axios.post(
    baseUrl + `/active-loans/${loanid}/reject`,
    {},
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
  return response.data;
};

interface activeloanbody {
  loanAmount: number;
  duration: number;
  interestRate: number;
  type: string;
}

type Form = {
  loanAmount: number;
  duration: number;
  interestRate: number;
  type: string;
};

const activeloans = async (token: string, data: Form) => {
  const response = await axios.post(
    baseUrl + "/active-loans",
    JSON.stringify(data),
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
  return response.data.data;
};

export {
  activeloansall,
  activeloansdetaildata,
  activeloansmyloans,
  activeloansid,
  activeloansidapprove,
  activeloansidreject,
  activeloans,
};

import axios from "axios";

const activeloansall = async (token: string) => {
  const response = await axios.get(
    "https://bank-dashboard-1tst.onrender.com/active-loans/all",
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
  const response = await axios.get(
    "https://bank-dashboard-1tst.onrender.com/active-loans/detail-data",
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
  return response.data.data;
};

const activeloansmyloans = async (token: string) => {
  const response = await axios.get(
    "https://bank-dashboard-1tst.onrender.com/active-loans/my-loans",
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
  return response.data.data;
};

const activeloansid = async (loanid:string,token: string) => {
  const response = await axios.get(
    `https://bank-dashboard-1tst.onrender.com/active-loans/${loanid}`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
  return response.data.data;
};

const activeloansidapprove = async (loanid:string,token: string) => {
  const respons7 = await axios.post(
    `https://bank-dashboard-1tst.onrender.com/active-loans/${loanid}/approve`,
    {},
    {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    }
  );
};

const activeloansidreject = async (loanid:string,token: string) => {
  const response = await axios.post(
    `https://bank-dashboard-1tst.onrender.com/active-loans/${loanid}/reject`,
    {},
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
  return response.data;
};

interface activeloanbody  {
  loanAmount: number;
  duration: number;
  interestRate: number;
  type: string;
}

const activeloans = async ( token: string) => {
  const response = await axios.post(
    "https://bank-dashboard-1tst.onrender.com/active-loans",
    JSON.stringify({
      loanAmount: 1,
      duration: 1,
      interestRate: 1,
      type: "personal",
    }),
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

import React from "react";
import Center from "./Center";
import Bottom from "./Bottom";
import fetchData from "@/app/Services/api/fetchData";

const Container = async () => {
  const endpoint = "https://bank-dashboard-6acc.onrender.com/cards";
  const accessToken = process.env.NAHOM_TOKEN as string;
  const res = await fetchData(endpoint, accessToken);
  console.log(res, "cards Data");
  return (
    <section className="w-full flex flex-col grow gap-6 p-8">
      <Center />
      <Bottom />
    </section>
  );
};

export default Container;

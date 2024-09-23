import React from "react";
import { redirect } from "next/navigation";

const page = () => {
  redirect("/dashboard");
  // return <div>page</div>;
};

export default page;

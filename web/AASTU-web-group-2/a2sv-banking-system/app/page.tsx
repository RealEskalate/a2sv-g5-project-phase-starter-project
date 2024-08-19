import { redirect } from "next/navigation";
import { getServerSession } from "next-auth";
import { options } from "./api/auth/[...nextauth]/options";
const page = async () => {
  const session = await getServerSession(options);

  if (!session) {
    console.log("NO SESSION")
    redirect("/api/auth/signin");

  } else {
    console.log(session)
    redirect("./dashboard");
  }
};

export default page;

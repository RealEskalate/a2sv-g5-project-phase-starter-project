'use client';
import { toggleSidebar } from "@/lib/redux/slices/menuSlice";
import { AppDispatch } from "@/lib/redux/store";
import { useDispatch } from "react-redux";
import Page from "./dashboard/Page";
import ServicePage from "./services/ServicePage";


export default function Home() {
  const dispatch: AppDispatch = useDispatch();

  const handleBurgerClick = () => {
    dispatch(toggleSidebar());
  };
  return (
    <div className="w-full" onClick={handleBurgerClick}>
        <Page/>
    </div>
  );
}

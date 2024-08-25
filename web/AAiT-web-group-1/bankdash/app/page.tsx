'use client';
import { toggleSidebar } from "@/lib/redux/slices/menuSlice";
import { AppDispatch } from "@/lib/redux/store";
import { useDispatch } from "react-redux";
import DashboardPage from "./components/Dashboard/DashboardPage";

export default function Home() {
  const dispatch: AppDispatch = useDispatch();

  const handleBurgerClick = () => {
    dispatch(toggleSidebar());
  };
  return (
    <div className="w-full" onClick={handleBurgerClick}>
        <DashboardPage/>
    </div>
  );
}

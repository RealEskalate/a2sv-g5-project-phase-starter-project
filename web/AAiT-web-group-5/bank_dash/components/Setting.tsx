import { CiSettings } from "react-icons/ci";
import { IoMdNotificationsOutline } from "react-icons/io";
import { LuCalendarCheck } from "react-icons/lu";

export default function Setting() {
  return (
    <div className="flex gap-3">
      <button
        type="button"
        className="rounded-full   p-1 w-10 h-10   bg-slate-200 flex justify-center items-center"
      >
        <CiSettings className="text-xl text-slate-700" />
      </button>

      <button
        type="button"
        className="rounded-full   p-1 w-10 h-10   bg-slate-200 flex justify-center items-center"
      >
        <IoMdNotificationsOutline className="text-xl text-red-400" />
      </button>
      
    </div>
  );
}

import { CiSearch } from "react-icons/ci";

export default function Search() {
  return (
    <div className="flex rounded-full   overflow-hidden h-10 mx-auto font-[sans-serif] text-black">
      <button
        type="button"
        className="flex items-center bg-slate-100 justify-center px-6"
      >
        <CiSearch className="text-slate-700 text-xl" />
      </button>
      <input
        placeholder="Search..."
        className="w-full outline-none bg-slate-100 text-sm px-5 py-3"
      />
    </div>
  );
}

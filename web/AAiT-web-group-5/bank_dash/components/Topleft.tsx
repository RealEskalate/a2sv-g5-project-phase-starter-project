import Search from "./Search";
import Seting from "./Seting";

export default function Topleft() {
  return (
    <div className="flex gap-3 mr-5 mt-3 w-1/3">
      <Search />
      <Seting />
    </div>
  );
}

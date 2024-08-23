import { Profile } from "./profile";
import Search from "./Search";
import Setting from "./Setting";

export default function Topright() {
  return (
    <div className="flex gap-3 mr-5 mt-3 w-1/3">
      <Search />
      <Setting />
      <Profile />
    </div>
  );
}

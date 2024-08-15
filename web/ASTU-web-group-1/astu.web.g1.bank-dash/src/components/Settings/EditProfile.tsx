import EditProfileForm from "../Form/EditProfileForm";
import ProfileAvatar from "./ProfileAvatar";

const EditProfile = () => {
  return (
    <div className="flex flex-col md:flex-row">
      <ProfileAvatar />
      <EditProfileForm />
    </div>
  );
};

export default EditProfile;

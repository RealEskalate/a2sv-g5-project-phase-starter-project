
import '../../../../core/constants/colors.dart';
import '../../../../core/constants/profile_photo.dart';
import '../../../auth/presentation/pages/pages.dart';

Widget showUser({  VoidCallback? onClicked}){
  return InkWell(
    onTap: () {
      onClicked ?? () {};
    },
    child: CircleAvatar(
      minRadius: 30,
      maxRadius: 40,
      backgroundColor: ChatColors.getRandomColor(),
    child: Image(
      image: AssetImage(ChatUserProfilePhoto.getRandomPhoto())),
    ),
  );
}
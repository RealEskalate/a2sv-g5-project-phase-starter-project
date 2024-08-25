import '../../../../core/constants/colors.dart';
import '../../../../core/constants/profile_photo.dart';
import '../../../auth/presentation/pages/pages.dart';

Widget showSentMessage(String text){
  return Column(
    crossAxisAlignment: CrossAxisAlignment.center,
    children: [

      Row(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          const Text('You',style: TextStyle(fontWeight: FontWeight.bold),),
          const SizedBox(width: 10,),
          CircleAvatar(
            backgroundColor: ChatColors.profileColors[0],
            child: Image(image: AssetImage(ChatUserProfilePhoto.photos[0])),
          )
          ],
      ),
      Container(
        padding: const EdgeInsets.all(10),
        decoration: BoxDecoration(
          color: ChatColors.mediumBlueColor,
          borderRadius: const  BorderRadius.only(topLeft: Radius.circular(10),bottomLeft: Radius.circular(10),bottomRight: Radius.circular(10))
        ),
        child: Text(text,style: const TextStyle(color: Colors.white),),
      ),
      const Text('03:25 AM')
    ],
  );
}
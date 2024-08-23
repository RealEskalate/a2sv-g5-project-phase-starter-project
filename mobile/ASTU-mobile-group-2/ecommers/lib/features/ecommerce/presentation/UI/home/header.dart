import 'package:date_formatter/date_formatter.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../../core/Colors/colors.dart';
import '../../../../../core/Text_Style/text_style.dart';
import '../../state/user_states/login_user_states_bloc.dart';
import '../../state/user_states/login_user_states_state.dart';
import 'profile/profile_page.dart';

class HeaderPart extends StatelessWidget {
  const HeaderPart({super.key});

  @override
  Widget build(BuildContext context) {
    String formattedDate = DateFormatter.formatDateTime(
      dateTime: DateTime.now(),
      outputFormat: 'dd-MM-yyyy',
    );
    return Container(
      padding: const EdgeInsets.only(top: 10),
      child: Row(
        children: [
          GestureDetector(
            onTap: () {
              showModalBottomSheet(
                  context: context,
                  builder: (BuildContext context) {
                    return const ProfilePage();
                  });
            },
            child: Container(
              width: 50,
              height: 50,
              decoration: BoxDecoration(
                  color: const Color.fromARGB(255, 202, 201, 201),
                  borderRadius: BorderRadius.circular(10)),
            ),
          ),
          const SizedBox(
            width: 10,
          ),
          SizedBox(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                SizedBox(
                  child: TextStyles(
                      text: formattedDate, fontColor: smallText, fontSizes: 12),
                ),
                BlocBuilder<LoginUserStatesBloc, LoginUserStates>(
                  builder: (context, user) {
                    final name = user is ProfileDetailState ? user.name : '';
                    return SizedBox(
                      child: Row(
                        children: [
                          TextStyles(
                            text: 'Hello, ',
                            fontColor: mainText,
                            fontSizes: 16,
                            fontWeight: FontWeight.w300,
                          ),
                          TextStyles(
                            text: name,
                            fontColor: mainText,
                            fontSizes: 20,
                            fontWeight: FontWeight.bold,
                            fontFamily: 'CaveatBrush',
                          ),
                        ],
                      ),
                    );
                  },
                )
              ],
            ),
          ),
          const Spacer(),
          Container(
            width: 30,
            height: 30,
            decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(5),
                color: const Color.fromARGB(255, 248, 247, 247),
                border: const Border(
                  top: BorderSide(color: Colors.grey),
                  right: BorderSide(color: Colors.grey),
                  left: BorderSide(color: Colors.grey),
                  bottom: BorderSide(color: Colors.grey),
                )),
            child: const Icon(
              Icons.notifications_none_sharp,
            ),
          )
        ],
      ),
    );
  }
}

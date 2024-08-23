import 'package:flutter/material.dart';

import '../../../../core/themes/themes.dart';

class UserInfo extends StatelessWidget {
  final String day;
  final String userName;
  final VoidCallback iconPres;
  const UserInfo(
      {super.key,
      required this.day,
      required this.userName,
      required this.iconPres});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(
        vertical: 10,
        horizontal: 30,
      ),
      child: Row(
        children: [
          Container(
            height: 50,
            width: 50,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(10),
              color: MyTheme.ecGrey,
              boxShadow: const [
                BoxShadow(
                  color: MyTheme.shadowColor,
                )
              ],
            ),
            // child: ClipRRect(
            //   child: Image.network(
            //     '',
            //     fit: BoxFit.fill,
            //   ),
            // ),
          ),
          Expanded(
            child: Padding(
              padding: const EdgeInsets.all(10),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    day,
                    style: const TextStyle(
                      color: MyTheme.ecTextGrey,
                      fontSize: 12,
                    ),
                  ),
                  RichText(
                    text: TextSpan(
                      text: 'Hellow, ',
                      style: const TextStyle(
                        color: Colors.black,
                      ),
                      children: [
                        TextSpan(
                          text: userName,
                          style: const TextStyle(
                            color: Colors.black,
                            fontWeight: FontWeight.bold,
                          ),
                        )
                      ],
                    ),
                  )
                ],
              ),
            ),
          ),
          Container(
            decoration: BoxDecoration(
                border: Border.all(color: MyTheme.ecGrey, width: 2),
                borderRadius: BorderRadius.circular(16)),
            child: IconButton(
              onPressed: iconPres,
              icon: const Icon(
                Icons.logout,
                color: Colors.black,
                size: 25,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

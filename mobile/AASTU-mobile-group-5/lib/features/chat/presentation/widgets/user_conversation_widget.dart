import 'package:flutter/material.dart';
import 'profile_widget.dart';

// ignore: camel_case_types
class user_conversation_widget extends StatelessWidget {
  final String name;
  const user_conversation_widget({
    super.key, required this.name,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(
        bottom: 25,
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            children: [
              const ProfileWidget(isOnline:true,iconUrl:'assets/images/Alex.png'),
              const SizedBox(
                width: 15,
              ),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    name,
                    style: const TextStyle(
                      color: Colors.black,
                      fontSize: 18,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                  const Text(
                    'How are you today?',
                    style: TextStyle(
                      color: Color.fromRGBO(121, 124, 123, 1),
                      fontSize: 12,
                      fontWeight: FontWeight.w400,
                    ),
                  ),
                ],
              ),
            ],
          ),
          const Column(
            children: [
              Text(
                '2 mins ago',
                style: TextStyle(
                  color: Color.fromRGBO(121, 124, 123, 1),
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ),
              SizedBox(
                height: 8,
              ),
              CircleAvatar(
                radius: 7,
                backgroundColor: Color.fromRGBO(73, 140, 240, 1),
                child: Text(
                  '4',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 10,
                    fontWeight: FontWeight.w700,
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}

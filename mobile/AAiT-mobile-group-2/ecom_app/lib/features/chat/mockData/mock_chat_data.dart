import 'package:ecom_app/features/chat/presentation/widgets/single_text.dart';
import 'package:flutter/material.dart';

final List<Widget> mock_chat_data = [
  SingleText(
    isMe: false,
    profile_pic: 'assets/profile_picture.png',
    name: 'Annei Ellison',
    text: 'Have a great working week!! Have a great working week!!',
    time: '09:25 AM',
  ),
  SingleText(
    isMe: true,
    profile_pic: 'assets/profile_picture.png',
    name: 'Annei Ellison',
    text: 'Have a great working week!! Have a great working week!!',
    time: '10:44 PM',
  ),
  SingleText(
    isMe: false,
    profile_pic: 'assets/profile_picture.png',
    name: 'Annei Ellison',
    image_content: 'assets/blue_portrait.jpg',
    time: '10:44 PM',
  ),
];

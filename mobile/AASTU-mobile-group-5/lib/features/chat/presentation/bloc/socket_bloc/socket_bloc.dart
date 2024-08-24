


import 'package:flutter/material.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class SocketBloc{



  final channel =  WebSocketChannel.connect(Uri.parse('https://g5-flutter-learning-path-be.onrender.com/'));
  

  // void dispose(List<String> arguments) async {
  
  // await Future.delayed(Duration(seconds: 5));
  // channel.sink.close();
}

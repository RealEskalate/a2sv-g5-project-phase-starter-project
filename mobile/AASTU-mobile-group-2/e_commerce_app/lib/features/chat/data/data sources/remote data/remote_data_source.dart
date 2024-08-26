import 'dart:convert';
import 'dart:developer';

import 'package:e_commerce_app/core/constants/constants.dart';
import 'package:e_commerce_app/core/services/auth_services.dart';

import '../../../../../core/failure/exception.dart';
import '../../models/chat_model.dart';
import '../../models/message_model.dart';
import 'remote_contrats.dart';
import 'package:http/http.dart' as http;
import 'stream_socket.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

import '../../../../../core/services/auth_services.dart';

class ChatRemoteDataSourceImpl extends ChatRemoteDataSource {
  final http.Client client;

  StreamSocket streamSocket = StreamSocket();
  // Replace 'YOUR_ACCESS_TOKEN' with the actual token you need

  ChatRemoteDataSourceImpl({
    required this.client,
  });

  @override
  Future<bool> deleteChatById(String chatId) async {
    try {
      final response =
          await client.delete(Uri.parse('${Urls.allChatsUrl}/$chatId'));

      if (response.statusCode != 200) {
        print(response.body);
        // throw ServerException();
        return false;
      } else {
        return true;
      }
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  @override
  Future<List<MessageModel>> getMessagesById(String chatId) async {
    // final socket = IO.io(
    //   Urls.socketUrl,
    //   IO.OptionBuilder()
    //       .setTransports(['websocket'])
    //       .disableAutoConnect()
    //       .setExtraHeaders({
    //         'Authorization': 'Bearer ${await AuthServices.getToken()}',
    //       })
    //       .build(),
    // );
    // streamSocket.dispose();
    // streamSocket = StreamSocket();
    print(chatId);
    client.get(Uri.parse('${Urls.allChatsUrl}/$chatId/messages'), headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ${await AuthServices.getToken()}',
    }).then((response) {
      if (response.statusCode == 200) {
        final List<dynamic> messages = jsonDecode(response.body)['data'];
        print(messages);
        return messages;
        // for (var message in messages) {
        //   streamSocket.addResponse(MessageModel.fromJson(message));
        // }
      } else {
        print(response.body);
        throw ServerException();
      }
    });
    throw ServerException();
  }

  //   socket.connect();

  //   socket.onConnect((_) {
  //     log('Connected to the socket server');
  //   });

  //   socket.onDisconnect((_) {
  //     log('Disconnected from the socket server');
  //   });

  //   socket.on('message:delivered', (data) {
  //     MessageModel message = MessageModel.fromJson(data);
  //     streamSocket.addResponse(message);
  //   });

  //   socket.on('message:received', (data) {
  //     MessageModel message = MessageModel.fromJson(data);
  //     streamSocket.addResponse(message);
  //   });

  //   return streamSocket.getResponse;
  // }

  @override
  Future<ChatModel> createChatById(String sellerId) async {
    try {
      // final response = await client.post(_baseUrl, {
      //   'userId': receiver.id,
      // });
      final response = await client.post(Uri.parse(Urls.allChatsUrl), body: {
        'userId': sellerId,
      }
      ,
      headers: {
      'Authorization': 'Bearer ${await AuthServices.getToken()}',
    }
      );

      if (response.statusCode == 201) {
        return ChatModel.fromJson(jsonDecode(response.body)['data']);
      } else {
        print(response.body);
        throw ServerException();
      }
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  @override
  Future<List<ChatModel>> getAllChat() async {
    try {
      final accessToken = AuthServices.getToken();
      print(accessToken);
      final response = await client.get(Uri.parse(Urls.allChatsUrl), headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ${await AuthServices.getToken()}',
      });
      streamSocket.dispose();
      streamSocket = StreamSocket();
      if (response.statusCode == 200) {
        print(response.body);
        final List<dynamic> chats = jsonDecode(response.body)['data'];
        return chats.map((e) => ChatModel.fromJson(e)).toList();
      } else {
        print(response.body);
        throw ServerException();
      }
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  @override
  Future<bool> sendMessage(String chatId, String message, String type) async {
    final socket = IO.io(
      Urls.socketUrl,
      IO.OptionBuilder()
          .setTransports(['websocket'])
          .disableAutoConnect()
          .setExtraHeaders({
            'Authorization': 'Bearer ${await AuthServices.getToken()}',
          })
          .build(),
    );
    print(chatId);
    print(message);
    print(type);
    try {
      socket.emit('message:send', {
        'chatId': chatId,
        'content': message,
        'type': type,
      });
      return true;
    } catch (e) {
      print(e);
      throw ServerException();
    }
  }

  
}








// void main() async{
// // final a = await AuthServices.getToken();
//   // Replace 'YOUR_ACCESS_TOKEN' with the actual token you need
//   final socket = IO.io(
//     'https://g5-flutter-learning-path-be.onrender.com/',
//     IO.OptionBuilder()
//       .setTransports(['websocket']) // Use WebSocket transport
//       .disableAutoConnect() // Disable auto-connect
//       .setExtraHeaders({
//         'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1vb29vYUBnbWFpbC5jb20iLCJzdWIiOiI2NmM0NGZkODYxOThmMTUwZTY0M2M4MjciLCJpYXQiOjE3MjQ2MzA5NzksImV4cCI6MTcyNTA2Mjk3OX0.ISII2Ra7ykjBg4-NUHgpw6P72O9S5slDiCbHmJM_kGs', // Add token to headers
//       })
//       .build(),
//   );

//   // Define a function to handle connection events
//   void onConnect(_) {
//     print('Connected');
    
//     // Send a test message once connected
//     socket.emit('message:send', {
//       "chatId": "66c837f2b7068ee15142f66a",
//       "content": "Hello, Mr. 00000",
//       "type": "text",
//     });
//     print('Message sent');
//   }

//   // Define a function to handle disconnection events
//   void onDisconnect(_) => print('Disconnected');

//   // Define a function to handle connection errors
//   void onError(error) {
//     print('Connection error: $error');
//     // Handle connection error logic here
//   }

//   // Define a function to handle 'message:delivered' events
//   void onMessageDelivered(data) {
//     print('Message delivered: $data');
//     // Handle message delivery logic here
//   }

//   // Define a function to handle 'message:received' events
//   void onMessageReceived(data) {
//     print('Message received: $data');
//     // Handle message reception logic here
//   }

//   // Set up event listeners
//   socket.onConnect(onConnect);
//   socket.onDisconnect(onDisconnect);
//   socket.onError(onError); // Listen for connection errors
//   socket.on('message:delivered', onMessageDelivered);
//   socket.on('message:received', onMessageReceived);

//   // Manually connect to the server
//   socket.connect();
// }

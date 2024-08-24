import 'package:dartz/dartz.dart';
import 'package:flutter/material.dart';
import 'core/error/failure.dart';
import 'features/chat_feature/chat/data_layer/model/chat_model.dart';
import 'features/chat_feature/chat/domain/entity/chat.dart';
import 'features/chat_feature/chat/domain/usecase/get_all_chat_history.dart';

class TempDart extends StatelessWidget {
  final GetAllChatHistory chatHistory;

  TempDart({required this.chatHistory, super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Chat History'),
      ),
      body: Center(
        child: StreamBuilder<Either<Failure, List<ChatEntity>>>(
          stream: chatHistory.call(),
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return CircularProgressIndicator(); 
            }

            if (snapshot.hasError) {
              return Text('Error: ${snapshot.error}'); 
            }

            if (snapshot.hasData) {
              print(snapshot.data);
            }

            return Text('No data available'); // Fallback UI
          },
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          print("Button pressed");
        },
        child: Icon(Icons.refresh),
      ),
    );
  }
}

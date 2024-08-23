import 'package:equatable/equatable.dart';

import '../../../../auth/domain/entities/user.dart';

class Message extends Equatable {
  final UserEntity sender;
  final String chatId;
  final String type;
  final String content;
 

  Message(
      {required this.sender,
      required this.chatId,
      required this.type,
      required this.content});
  @override
  List<Object?> get props => throw UnimplementedError();
}

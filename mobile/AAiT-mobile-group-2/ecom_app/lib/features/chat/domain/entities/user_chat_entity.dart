import 'package:equatable/equatable.dart';

class UserChatEntity extends Equatable {
  final String ChatID;
  final String name;

  UserChatEntity({required this.ChatID, required this.name});

  @override
  List<Object?> get props => [ChatID, name];
}

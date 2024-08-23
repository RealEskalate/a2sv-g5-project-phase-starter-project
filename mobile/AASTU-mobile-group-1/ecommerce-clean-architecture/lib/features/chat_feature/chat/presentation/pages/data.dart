import '../../../../auth/domain/entities/user.dart';
import '../../domain/entity/message.dart';

List<Message> messages = [
  Message(
    sender: UserEntity(
        id: "002", name: "dina", email: "dina@gmail.com", password: "12345678"),
    chatId: '002',
    type: 'text',
    content: 'Hello',
  ),
  Message(
    sender: UserEntity(
        id: "001", name: "dina", email: "dina@gmail.com", password: "12345678"),
    chatId: '002',
    type: 'text',
    content: 'Hello',
  ),
  Message(
    sender: UserEntity(
        id: "002", name: "dina", email: "dina@gmail.com", password: "12345678"),
    chatId: '002',
    type: 'text',
    content: 'Hello',
  ),
  Message(
    sender: UserEntity(
        id: "001", name: "dina", email: "dina@gmail.com", password: "12345678"),
    chatId: '002',
    type: 'text',
    content: 'Hello',
  ),
];

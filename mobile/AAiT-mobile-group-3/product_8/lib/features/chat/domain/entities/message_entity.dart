import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_data_entity.dart';

class MessageEntity extends Equatable {
  final String id;
  final UserDataEntity sender;
  final String content;
  final String type;

  const MessageEntity({
    required this.id,
    required this.sender,
    required this.content,
    required this.type,
  });
  @override
  List<Object?> get props => [id, sender, content, type];
}

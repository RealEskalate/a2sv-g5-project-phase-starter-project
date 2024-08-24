import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_data_entity.dart';

class MessageEntity extends Equatable {
  final String _id;
  final UserDataEntity sender;
  final String content;
  final String type;

  const MessageEntity(
    this._id, {
    required this.sender,
    required this.content,
    required this.type,
  });
  @override
  List<Object?> get props => [_id, sender, content, type];
}

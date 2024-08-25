import 'package:equatable/equatable.dart';

import '../../../product/domain/entitity/user.dart';

class Chat extends Equatable {
  final String id;
  final User sender;
  final User receiver;

  const Chat({
    required this.id,
    required this.sender,
    required this.receiver,
  }) : super();

  @override
  List<Object?> get props => [id, sender, receiver];
}

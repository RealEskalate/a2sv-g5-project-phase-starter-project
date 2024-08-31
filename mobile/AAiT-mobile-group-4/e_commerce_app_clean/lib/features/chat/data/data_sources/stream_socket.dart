import 'dart:async';

import '../models/message_model.dart';

class StreamSocket {
  final socketResponse = StreamController<MessageModel>();

  void Function(MessageModel) get addResponse => socketResponse.sink.add;
  
  Stream<MessageModel> get getResponse => socketResponse.stream;

}
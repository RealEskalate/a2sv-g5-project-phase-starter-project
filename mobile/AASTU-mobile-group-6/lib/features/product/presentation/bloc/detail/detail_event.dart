import 'package:equatable/equatable.dart';


abstract class DetailEvent extends Equatable {
  const DetailEvent();
}

class DeleteProductEvent extends DetailEvent {
  final String id;
  DeleteProductEvent(this.id);
  @override
  List<Object> get props => [id];
}


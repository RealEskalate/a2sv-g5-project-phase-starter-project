

import 'package:equatable/equatable.dart';

class Failur  extends Equatable{
  final String message;
  const Failur ({
    required this.message
  });


  @override
  List<Object?> get props => [
    message
  ];
}


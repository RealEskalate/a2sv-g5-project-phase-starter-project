import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:equatable/equatable.dart';


abstract class AddEvent extends Equatable {
  const AddEvent();
}

class AddProductEvent extends AddEvent {
  final ProductEntity product;
  const AddProductEvent({required this.product});
  @override
  List<Object> get props => [product];

}


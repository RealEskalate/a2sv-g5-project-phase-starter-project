import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
import 'package:equatable/equatable.dart';


abstract class UpdateEvent extends Equatable {
  const UpdateEvent();
}

class UpdateProductEvent extends UpdateEvent {
  final ProductModel product;
  const UpdateProductEvent({required this.product});
  @override
  List<Object> get props => [product];

}

part of 'update_product_bloc.dart';

class ProductUpdatedEvent {
  ProductUpdatedEvent();
}

class ProductUpdated extends ProductUpdatedEvent {
  ProductEntity product;
  ProductUpdated({required this.product});
}

class UpdateInitiated extends ProductUpdatedEvent {
  UpdateInitiated();
}

class ProductDeleted extends ProductUpdatedEvent {
   ProductEntity product;
  ProductDeleted({required this.product});
}

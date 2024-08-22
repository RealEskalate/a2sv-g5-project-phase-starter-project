part of 'update_product_bloc.dart';

abstract class UpdateProductState {}

class UpdateProductInitial extends UpdateProductState {
  UpdateProductInitial();
}

class Update extends UpdateProductState {
  Update();
}

class UpdateProductLoading extends UpdateProductState {
  UpdateProductLoading();
}

class UpdateProductFail extends UpdateProductState {
  UpdateProductFail();
}

class UpdateProductSuccess extends UpdateProductState {
  UpdateProductSuccess();
}

import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:ecommerce/core/import/import_file.dart';
import 'package:meta/meta.dart';

part 'detail_product_event.dart';
part 'detail_product_state.dart';

class DetailProductBloc extends Bloc<DetailProductEvent, DetailProductState> {
  ProductEntity product;
  DetailProductBloc({
    required this.product,
  }) : super(DetailProductState(product: product)) {
    on<DetailProductEvent>(_detailPage);
    
  }

  FutureOr<void> _detailPage(DetailProductEvent event, Emitter<DetailProductState> emit) {
  }
}

/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package assets

import (
	"fmt"
	"strings"
)

type SymbolInfo interface {

	//CurveType 曲线类型
	CurveType() uint32

	//FullName 币种全名
	FullName() string

	//Symbol 币种标识
	Symbol() string

}

// GetSymbolInfo 获取资产的币种信息
func GetSymbolInfo(symbol string) (SymbolInfo, error) {
	manager, ok := Managers[strings.ToLower(symbol)].(SymbolInfo)
	if !ok {
		return nil, fmt.Errorf("assets: %s is not support", symbol)
	}
	return manager, nil
}

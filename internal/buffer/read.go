package buffer

import "io"

// ReadFullLimited 读取至多 limit 字节。
func ReadFullLimited(r io.Reader, limit int) ([]byte, error) {
	if limit <= 0 {
		return nil, io.ErrUnexpectedEOF
	}
	buf := make([]byte, limit)
	n, err := io.ReadFull(r, buf)
	if err == io.ErrUnexpectedEOF || err == io.EOF {
		return CloneBytes(buf[:n]), nil
	}
	if err != nil {
		return nil, err
	}
	return buf, nil
}
